import { Component, Show, createEffect, createMemo, createSignal, onMount } from "solid-js";

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

const getNewSimulation = async () => {
  const response = await fetch("http://localhost:3000/api/2021");
  const data = await response.json();
  return data;
}


const App: Component = () => {
  const [playedTournament, setPlayedTournament] = createSignal<PlayedTournament>({
    id: 0,
    name: "",
    groups: [],
    teams: [],
    matches: [],
    matchResults: [],
    groupRankings: [],
    koMatches: []
  });

  const groupPhaseGroups = createMemo(() => {
    return playedTournament().groups.filter((group) => group.groupType === "group_phase");
  });

  const newSimulation = async () => {
    const data = await getNewSimulation() as PlayedTournament;
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
          <p class="py-4">EM soccer tournament simulator 2021</p>
        </div>

        <button onClick={newSimulation} class="bg-white text-black rounded-md px-4 py-2 mr-4">Run simulation</button>
      </header>
      <AllGroupMatches groups={groupPhaseGroups()} matchResults={playedTournament().matchResults} />
      <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">KO phase</h1>
      {/*
      <div class="flex flex-row flex-wrap">
        <KoRound matches={tournamentOutcome().sixteen} name="Round of 16" />
        <KoRound matches={tournamentOutcome().eight} name="Round of 8" />
        <KoRound matches={tournamentOutcome().four} name="Round of 4" />
      </div>
      <div class="flex justify-around">
        <KoRound matches={tournamentOutcome().final} name="Final" />
      </div>
      <Show when={tournamentOutcome().final.length > 0} fallback={<h1 class="p-4 text-xl bg-sky-800 bg-opacity-25">No Winner yet</h1>}>
        <h1 class="p-4 text-xl bg-sky-800 bg-opacity-25 text-center">Winner: {tournamentOutcome().final[0].winner.name}</h1>
      </Show> */}
    </div>
  );
};

export default App;
