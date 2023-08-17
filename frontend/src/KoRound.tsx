import { Component, For, Show } from "solid-js";

const KoRound: Component = (props: {matches: Array<any>, name: string}) => {
  return (
    <section>
      <h1>{props.name}</h1>
      <For each={props.matches}>{(match, i) =>
        <div class="flex">
          <div>
            {match.team1.name} - {match.team2.name}: {match.goalsTeam1} : {match.goalsTeam2}&nbsp;
            <Show when={match.goalsTeam1===match.goalsTeam2}>
              (a.P {match.penaltyScoreTeam1} - {match.penaltyScoreTeam2})
            </Show>
          </div>
        </div>}
      </For>
    </section>
  )
}

export default KoRound
