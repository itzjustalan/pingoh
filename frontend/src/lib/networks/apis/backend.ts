import axios, { type AxiosInstance } from 'axios';
import { env } from '../../../env';
import { log } from '../../logger';
import { authStore } from '../../stores/auth';
export { type AxiosInstance } from 'axios';

const backendApi: AxiosInstance = axios.create({
	baseURL: `http://${env.dev ? 'localhost:3000' : window.location.host}/api`
});

backendApi.defaults.headers.common['Content-Type'] = 'application/json';
backendApi.interceptors.request.use(
	(config) => {
    config.headers.Authorization = `Bearer ${authStore.getState().user?.access_token}`;
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
			config.config.method?.toUpperCase() ?? '-',
			config.config.url ?? '-',
			config.data
		);
		return config;
	},
	(error) => {
		log.cl_res(
			error.response?.status ?? 0,
			error.response?.statusText ?? error.message,
			error.config?.method?.toUpperCase() ?? '-',
			error.config?.url ?? '-',
			error.response?.data
		);
	}
);

export default backendApi;
