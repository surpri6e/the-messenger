import { create } from "zustand";

interface IUseRegistrationDataStoreActions {
  setNewEmail: (email: string) => void;
  setNewPassword: (password: string) => void;
  setNewUsername: (nickname: string) => void;

  setDefault: () => void;
}

interface IUseRegistrationDataStore {
  email: string;
  password: string;
  username: string;

  actions: IUseRegistrationDataStoreActions;
}

export const useRegistrationDataStore = create<IUseRegistrationDataStore>(
  (set) => ({
    email: "",
    username: "",
    password: "",

    actions: {
      setNewEmail: (email: string) => set((state) => ({ ...state, email })),
      setNewUsername: (username: string) =>
        set((state) => ({ ...state, username })),
      setNewPassword: (password: string) =>
        set((state) => ({ ...state, password })),

      setDefault: () => set({ email: "", username: "", password: "" }),
    },
  }),
);
