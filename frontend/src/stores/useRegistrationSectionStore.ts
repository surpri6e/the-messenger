import { create } from "zustand";

type TSectionType = "registration" | "auth" | "restore";

interface IUseRegistrationSectionStoreActions {
  setRegistrationType: () => void;
  setAuthType: () => void;
  setRestoreType: () => void;
}

interface IUseRegistrationSectionStore {
  sectionType: TSectionType;
  actions: IUseRegistrationSectionStoreActions;
}

export const useRegistrationSectionStore = create<IUseRegistrationSectionStore>(
  (set) => ({
    sectionType: "registration",

    actions: {
      setRegistrationType: () => set({ sectionType: "registration" }),
      setAuthType: () => set({ sectionType: "auth" }),
      setRestoreType: () => set({ sectionType: "restore" }),
    },
  }),
);
