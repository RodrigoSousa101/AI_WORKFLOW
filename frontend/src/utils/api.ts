import axios from 'axios';

// Criação da instância com baseURL e withCredentials
const api = axios.create({
  baseURL: 'http://localhost:8080/api',
  withCredentials: true, // IMPORTANTE para enviar cookies como refresh_token
});

// Função de logout segura
function logout() {
  localStorage.removeItem('access_token');
  document.cookie = "refresh_token=; Max-Age=0; path=/;";
  window.location.href = '/';
}

let isRefreshing = false;
let refreshSubscribers: Function[] = [];

function onRefreshed(token: string) {
  refreshSubscribers.forEach((cb) => cb(token));
  refreshSubscribers = [];
}

// Adiciona o access_token no header Authorization
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Intercepta respostas com erro (401)
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (!originalRequest) {
      return Promise.reject(error);
    }

    // 🔴 Se o refresh falhou (ex: refresh expirado ou ausente), fazer logout direto
    if (
      error.response?.status === 401 &&
      originalRequest.url.includes('/auth/refresh')
    ) {
      logout();
      return Promise.reject(error);
    }

    // Se já tentamos e falhou, forçar logout
    if (originalRequest._retry) {
      logout();
      return Promise.reject(error);
    }

    // Se deu 401, tentar refresh
    if (error.response?.status === 401) {
      originalRequest._retry = true;

      if (isRefreshing) {
        return new Promise((resolve, reject) => {
          refreshSubscribers.push((token: string) => {
            originalRequest.headers.Authorization = `Bearer ${token}`;
            resolve(api(originalRequest));
          });
        });
      }

      isRefreshing = true;

      try {
        const refreshResponse = await api.post('/auth/refresh');
        const newAccessToken = refreshResponse.data.access;
        localStorage.setItem('access_token', newAccessToken);
        originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
        onRefreshed(newAccessToken);
        isRefreshing = false;
        return api(originalRequest);
      } catch (refreshError) {
        isRefreshing = false;
        logout();
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

export default api;
export { logout };
