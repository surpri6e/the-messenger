import { create } from "zustand";

interface IUseAuthDataStoreActions {
  setNewEmail: (email: string) => void;
  setNewPassword: (password: string) => void;

  setDefault: () => void;
}

interface IUseAuthDataStore {
  email: string;
  password: string;

  actions: IUseAuthDataStoreActions;
}

export const useAuthDataStore = create<IUseAuthDataStore>((set) => ({
  email: "",
  password: "",

  actions: {
    setNewEmail: (email: string) => set((state) => ({ ...state, email })),
    setNewPassword: (password: string) =>
      set((state) => ({ ...state, password })),

    setDefault: () => set({ email: "", password: "" }),
  },
}));
