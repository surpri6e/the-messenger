import type { IChat, ICommunication, IGroup } from "@appTypes/ICommunication";
import { create } from "zustand";
import { createJSONStorage, persist } from "zustand/middleware";

interface IUseCurrentCommunicationStoreActions {
  setCurrentChat: (chat: IChat) => void;
  setCurrentGroup: (group: IGroup) => void;
}

interface IUseCurrentCommunicationStore {
  currentCommunication: ICommunication;

  actions: IUseCurrentCommunicationStoreActions;
}

export const useCurrentCommunicationStore =
  create<IUseCurrentCommunicationStore>()(
    persist(
      (set) => ({
        currentCommunication: {
          currentChat: null,
          currentGroup: null,
          type: null,
        },

        actions: {
          setCurrentChat: (chat: IChat) =>
            set((state) => ({
              currentCommunication: {
                type: "chat",
                currentChat: chat,
                currentGroup: state.currentCommunication.currentGroup,
              },
            })),

          setCurrentGroup: (group: IGroup) =>
            set((state) => ({
              currentCommunication: {
                type: "group",
                currentChat: state.currentCommunication.currentChat,
                currentGroup: group,
              },
            })),
        },
      }),

      {
        name: "currentCommunication",
        storage: createJSONStorage(() => sessionStorage),
        partialize: (state) => ({
          currentCommunication: state.currentCommunication,
        }),
      },
    ),
  );
