import {
  type SigninInput,
  type SignupInput,
  signinInputSchema,
  signupInputSchema,
} from "../models/inputs/auth";
import { type AuthedUser, authStore } from "../stores/auth";
import { decodeJwt } from "../utils";
import backendApi from "./apis/backend";

class AuthNetwork {
  accessTimeout: undefined | number;
  constructor() {
    const stored = JSON.parse(localStorage.getItem("auth") ?? "null");
    const token = stored?.state?.user?.access_token;
    if (token) this._autoRefresh(token);
  }

  signout = () => {
    clearInterval(this.accessTimeout);
    authStore.setState((state) => ({ ...state, user: undefined }));
    // invalidateAll();
  };

  signup = async (data: SignupInput): Promise<AuthedUser> => {
    signupInputSchema.parse(data);
    const res = await backendApi.post<AuthedUser>("/auth/signup", data);
    this._autoRefresh(res.data.access_token);
    authStore.setState((state) => ({ ...state, user: res.data }));
    return res.data;
  };

  signin = async (data: SigninInput): Promise<AuthedUser> => {
    signinInputSchema.parse(data);
    const res = await backendApi.post<AuthedUser>("/auth/signin", data);
    this._autoRefresh(res.data.access_token);
    authStore.setState((state) => ({ ...state, user: res.data }));
    return res.data;
  };

  refresh = async () => {
    const response = await backendApi.post<
      Pick<AuthedUser, "access_token" | "refresh_token">
    >("/auth/refresh", {
      token: authStore.getState().user?.refresh_token,
    });
    this._autoRefresh(response.data.access_token);
    authStore.setState((state) => {
      if (!state.user) return state;
      state.user.access_token = response.data.access_token ?? "";
      state.user.refresh_token = response.data.refresh_token ?? "";
      return state;
    });
    return response.data;
  };

  _autoRefresh = (accessToken: string) => {
    const decodedToken = decodeJwt(accessToken);
    clearTimeout(this.accessTimeout);
    this.accessTimeout = setTimeout(
      this.refresh,
      (decodedToken.exp - decodedToken.iat) * 1000,
    );
  };
}

export const authNetwork = new AuthNetwork();
