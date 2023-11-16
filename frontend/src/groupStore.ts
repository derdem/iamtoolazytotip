import { createStore } from "solid-js/store";

export enum Strength {
  Weak = 1,
  Medium = 2,
  Strong = 3
}

export interface Country {
  name: string
  strength: Strength
}

export interface GroupInStore {
  groupName: string
  countries: Country[],
  index: number
}

const createGroupStore = () => {
  return createStore<GroupInStore[]>([])
}

export const [groups, setGroups] = createGroupStore()

