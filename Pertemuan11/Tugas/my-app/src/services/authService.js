import axios from "axios";

const API_URL_LOGIN = "http://127.0.0.1:8088/login";
const API_URL_REGISTER = "http://127.0.0.1:8088/register";

export const login = async (username, password) => {
  const response = await axios.post(API_URL_LOGIN, { username, password });
  return response.data;
};

export const register = async (userData) => {
  const response = await axios.post(API_URL_REGISTER, userData);
  return response.data;
};