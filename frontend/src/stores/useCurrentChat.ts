import type { IChat } from "@appTypes/IChat";
import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

interface IUseCurrentChatStoreActions {
  setCurrentChat: (chat: IChat | null) => void;
}

interface IUseCurrentChatStore {
  currentChat: IChat | null;

  actions: IUseCurrentChatStoreActions;
}

export const useCurrentChatStore = create<IUseCurrentChatStore>()(
  persist(
    (set) => ({
      currentChat: null,

      actions: {
        setCurrentChat: (chat: IChat | null) =>
          set(() => ({ currentChat: chat })),
      },
    }),

    {
      name: "currentChat",
      storage: createJSONStorage(() => sessionStorage),
      partialize: (state) => ({
        currentChat: state.currentChat,
      }),
    },
  ),
);
