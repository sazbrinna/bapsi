import axios from "axios";
import { URL } from "@/axios/ApiBaseURL";

const API = axios.create({
  baseURL: URL.internet2,
  timeout: 5000,
});

// Add a request interceptor
API.interceptors.request.use(
  function (config) {
    // Get the token from local storage
    const token = localStorage.getItem("tokenAuth");

    // If token exists, add it to the Authorization header
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);

// Add a response interceptor
API.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    // Handle 401 Unauthorized error
    if (error.response.status === 401 || error.response.status === 403) {
      localStorage.clear();
      window.location.href = "/login";
    }
    return Promise.reject(error);
  }
);
export { API as axios };
