import { type AuthInput, authInputSchema } from '$lib/models/input/user';
import backendApi from './apis/backend';
import { decodeJwt } from '$lib/utils';
import { authedUser, type AuthedUser } from '$lib/stores/auth';
import { get } from 'svelte/store';

class AuthNetwork {
	accessTimeout: NodeJS.Timeout | undefined;

	// refresh = async () => await defaultApi.get('v1/api/auth/refresh');
	// signin = async (data: AuthData): Promise<AuthResponse> =>
	//   (await defaultApi.post("v1/api/auth/signin", authSchema.parse(data))).data;

	signout = () => {
		clearInterval(this.accessTimeout);
		authedUser.clear();
	};

	signup = async (data: AuthInput): Promise<AuthedUser> => {
		authInputSchema.parse(data);
		const res = await backendApi.post<AuthedUser>('/auth/signup', data);
		this._autoRefresh(res.data.access_token);
		authedUser.set(res.data);
		return res.data;
	};

	signin = async (data: AuthInput): Promise<AuthedUser> => {
		authInputSchema.parse(data);
		const res = await backendApi.post<AuthedUser>('/auth/signin', data);
		this._autoRefresh(res.data.access_token);
		authedUser.set(res.data);
		return res.data;
	};

	refresh = async (): Promise<AuthedUser> => {
		const response = await backendApi.post<AuthedUser>('/auth/refresh', {
			token: get(authedUser)?.refresh_token,
		});
		this._autoRefresh(response.data.access_token);
		authedUser.updateTokens(response.data);
		return response.data;
	};

	_autoRefresh = (accessToken: string) => {
		const decodedToken = decodeJwt(accessToken);
		clearTimeout(this.accessTimeout);
		this.accessTimeout = setTimeout(this.refresh, (decodedToken.exp - decodedToken.iat) * 1000);
	};
}

export const authNetwork = new AuthNetwork();
