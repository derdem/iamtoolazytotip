import { Component, For } from "solid-js";

const SingleGroupMatches: Component<{ matches: any[] }> = (props) => {
  return (
    <div>
      <h2 class="pb-2 text-lg">{props.matches[0].groupName}</h2>
      <For each={props.matches}>
        {(match, i) => (
          <div class="flex">
            <div class="whitespace-nowrap">
              {match.team1.name} - {match.team2.name} : {match.goalsTeam1} :{" "}
              {match.goalsTeam2}
            </div>
          </div>
        )}
      </For>
    </div>
  );
};

export default SingleGroupMatches;
