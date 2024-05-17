import { Component, For, Show, createEffect, createMemo, createSignal, onMount } from "solid-js";

import AllGroupMatches from "../AllGroupMatches";
import KoRound from "../KoRound";
import { A } from "@solidjs/router";
import { PlayedTournament } from "../types";

const sortGroupMatches = (groupMatches: any[]) => {
  const sortedMatches: any = {};
  groupMatches.forEach((match) => {
    if (!sortedMatches[match.groupName]) {
      sortedMatches[match.groupName] = [];
    }
    sortedMatches[match.groupName].push(match);
  });
  console.log(sortedMatches);

  const compareGroupName = (a: any, b: any) => {
    if (a < b) {
      return -1;
    }
    if (a > b) {
      return 1;
    }
    return 0;
  }

  const sortedGroupNames = Object.keys(sortedMatches).sort(compareGroupName);

  const sortedMatchesAsArray: any[] = [];
  sortedGroupNames.forEach((groupName) => {
    sortedMatchesAsArray.push(sortedMatches[groupName]);
  });

  return sortedMatchesAsArray;
};

const getNewSimulation = async (url: string) => {
  const response = await fetch(url);
  const data = await response.json();
  return data;
}


const App: (u: string, n: string) => Component = (url: string, name: string) => () => {
  const [playedTournament, setPlayedTournament] = createSignal<PlayedTournament>({
    id: 0,
    name: "",
    groups: [],
    teams: [],
    matches: [],
    matchResults: [],
    groupRankings: [],
    koMatches: [],
    winner: {},
  });

  const groupPhaseGroups = createMemo(() => {
    return playedTournament().groups.filter((group) => group.groupType === "group_phase");
  });
  const koPhaseGroups = createMemo(() => {
    return playedTournament().groups.filter((group) => group.groupType === "knockout_phase");
  });

  const filterMatchResults = (groupId: number) => {
    console.log(`groupId: ${groupId}`)
    return playedTournament().matchResults.filter(
      (matchResult) => matchResult.match.groupId === groupId
    );
  }

  const newSimulation = async () => {
    const data = await getNewSimulation(url) as PlayedTournament;
    console.log(data);
    setPlayedTournament(data);
  }

  onMount(async () => {
    newSimulation()
  });

  return (
    <div>
      <header class="bg-sky-800 text-center text-white flex justify-between items-center">
        <div class="flex text-2xl ml-4">
          <A href="/" class="no-underline"><i class="py-4 mr-4 fa-solid fa-house"></i></A>
          <p class="py-4">{name}</p>
        </div>

        <button onClick={newSimulation} class="bg-white text-black rounded-md px-4 py-2 mr-4">Run simulation</button>
      </header>
      <AllGroupMatches groups={groupPhaseGroups()} matchResults={playedTournament().matchResults} />
      <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">KO phase</h1>
      <For each={koPhaseGroups()}>
        {(group, i) => (
          <div class="p-4">
            <KoRound matchResults={filterMatchResults(group.id)} name={group.name} />
          </div>
        )}
      </For>
      <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25 text-center">Winner: {playedTournament().winner.name}</h1>
    </div>
  );
};

export default App;
