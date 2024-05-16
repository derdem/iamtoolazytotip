import { Component, Show } from "solid-js";

const TeamMatchKo: Component<{ [key: string]: any }> = (props) => {
  return (
    <>
      <div class="">{props.matchResult.match.team1.name}</div>
      <div class="px-4">&nbsp;-&nbsp;</div>
      <div class="">{props.matchResult.match.team2.name}</div>
      <div class="flex-grow pl-4">
        :&nbsp;&nbsp;&nbsp;
        {props.matchResult.team1Goals}&nbsp;-&nbsp;{props.matchResult.team2Goals}
        &nbsp;
        <Show when={props.matchResult.team1Goals === props.matchResult.team2Goals}>
          (a.P {props.matchResult.team1PenaltyGoals} - {props.matchResult.team2PenaltyGoals}
          )
        </Show>
      </div>
    </>
  );
};

export default TeamMatchKo;
