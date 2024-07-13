// import { PUBLIC_ADMIN_API_URL } from '$env/static/public';
import { dev } from '$app/environment';
import { log } from '$lib/logger';
import { auth } from '$lib/stores/auth';
import axios, { type AxiosInstance } from 'axios';
export { type AxiosInstance } from 'axios';

const backendApi: AxiosInstance = axios.create({
	baseURL: `http://${dev ? 'localhost:3000' : window.location.host}/api`
});

backendApi.defaults.headers.common['Content-Type'] = 'application/json';
backendApi.interceptors.request.use(
	(config) => {
		config.headers.Authorization = 'Bearer ' + auth.user?.access_token;
		log.cl_req(config.method ?? '-', config.url ?? '-', config.data);
		return config;
	},
	(error) => {
		log.cl_req(error.config.method ?? '-', error.config.url ?? '-', error.response.data);
	}
);
backendApi.interceptors.response.use(
	(config) => {
		log.cl_res(
			config.status,
			config.statusText,
			config.config.method ?? '-',
			config.config.url ?? '-',
			config.data
		);
		return config;
	},
	(error) => {
		log.cl_res(
			error.response.status,
			error.response.statusText,
			error.config.method ?? '-',
			error.config.url ?? '-',
			error.response.data
		);
	}
);

// defaultApi.interceptors.request.use(config => {
//   config.headers.Authorization = '';
//   return config;
// })

// defaultApi.interceptors.request.use(config => {
//   return config;
// })
// defaultApi.interceptors.response.use(config => {
//   log.request(config.status, config.request.method, '', 0);
//   return config;
// })

export default backendApi;
