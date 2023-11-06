import { Component, Show } from "solid-js";

const TeamMatchKo: Component<{ [key: string]: any }> = (props) => {
  return (
    <>
      <div class="">{props.match.team1.name}</div>
      <div class="px-4">&nbsp;-&nbsp;</div>
      <div class="">{props.match.team2.name}</div>
      <div class="flex-grow pl-4">
        :&nbsp;&nbsp;&nbsp;
        {props.match.goalsTeam1}&nbsp;-&nbsp;{props.match.goalsTeam2}
        &nbsp;
        <Show when={props.match.goalsTeam1 === props.match.goalsTeam2}>
          (a.P {props.match.penaltyScoreTeam1} - {props.match.penaltyScoreTeam2}
          )
        </Show>
      </div>
    </>
  );
};

export default TeamMatchKo;
