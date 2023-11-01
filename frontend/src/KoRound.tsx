import { Component, For, Show } from "solid-js";

const KoRound: Component<{ matches: Array<any>; name: string }> = (props) => {
  return (
    <section class="p-4 w-1/3">
      <h1 class="pb-2 text-lg">{props.name}</h1>
      <For each={props.matches}>
        {(match, i) => (
          <div class="flex">
            <div>
              {match.team1.name} - {match.team2.name}: {match.goalsTeam1} :{" "}
              {match.goalsTeam2}&nbsp;
              <Show when={match.goalsTeam1 === match.goalsTeam2}>
                (a.P {match.penaltyScoreTeam1} - {match.penaltyScoreTeam2})
              </Show>
            </div>
          </div>
        )}
      </For>
    </section>
  );
};

export default KoRound;
