import { create } from "zustand";
import { persist } from "zustand/middleware";
import type { UserModel } from "../models/db/user";

export interface AuthedUser extends UserModel {
  access_token: string;
  refresh_token: string;
}

type AuthStoreState = {
  user: undefined | AuthedUser;
};

export const authStore = create<AuthStoreState>()(
  persist(
    (set) => ({
      user: undefined,
      _set: (u: AuthedUser | undefined) =>
        set((state) => ({ ...state, user: u })),
      _updateTokens: (tokens: Partial<AuthedUser>) => set((state) => {
        if (!state.user) return state;
        state.user.access_token = tokens.access_token ?? '';
        state.user.refresh_token = tokens.refresh_token ?? '';
        return state;
      }),
    }),
    { name: "auth" },
  ),
);
