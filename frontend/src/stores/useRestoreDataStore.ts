import { create } from "zustand";

interface IUseRestoreDataStoreActions {
  setNewEmail: (email: string) => void;
  setNewPassword: (password: string) => void;
  setNewCode: (code: string) => void;

  setDefault: () => void;
}

interface IUseRestoreDataStore {
  email: string;
  password: string;
  code: string;

  actions: IUseRestoreDataStoreActions;
}

export const useRestoreDataStore = create<IUseRestoreDataStore>((set) => ({
  email: "",
  password: "",
  code: "",

  actions: {
    setNewEmail: (email: string) => set((state) => ({ ...state, email })),
    setNewPassword: (password: string) =>
      set((state) => ({ ...state, password })),
    setNewCode: (code: string) => set((state) => ({ ...state, code })),

    setDefault: () => set({ email: "", password: "" }),
  },
}));
