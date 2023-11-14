import { createStore } from "solid-js/store";

const createGroupStore = () => {
  return createStore<string[]>([])
}

export const [groups, setGroups] = createGroupStore()

