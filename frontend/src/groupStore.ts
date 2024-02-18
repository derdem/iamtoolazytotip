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

export interface MatchInStore {
  groupIndex: number
  matchIndex: number
  country1: string
  country2: string
}

const createGroupStore = () => {
  return createStore<GroupInStore[]>([])
}

const createMatchStore = () => {
  return createStore<MatchInStore[]>([])
}

export const createEmptyMatch: () => MatchInStore = () => (
  {groupIndex: 0, matchIndex: 0, country1: "", country2: ""}
)

export const [groups, setGroups] = createGroupStore()
export const [matches, setMatches] = createMatchStore()

