import type { IUser } from "@appTypes/IUser";
import { create } from "zustand";

interface IUseUserStoreActions {
  setUser: (user: IUser | null) => void;
}

interface IUseUserStore {
  user: IUser | null;
  userContacts: number[] | null;
  userEnemiesId: number[] | null;

  actions: IUseUserStoreActions;
}

export const useUserStore = create<IUseUserStore>((set) => ({
  user: null,
  userContacts: null,
  userEnemiesId: null,

  actions: {
    setUser: (user: IUser | null) => set((state) => ({ ...state, user })),
  },
}));
