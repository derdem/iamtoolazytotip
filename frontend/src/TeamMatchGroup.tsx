import { Component, Show } from "solid-js";

const TeamMatchGroup: Component<{[key: string]: any}> = (props) => {
  return (
    <>
      <div class="">{props.match.team1.name}</div>
      <div class="px-4">&nbsp;-&nbsp;</div>
      <div class="">{props.match.team2.name}</div>
      <div class="flex-grow pl-4">
        :&nbsp;&nbsp;&nbsp;
        {props.match.goalsTeam1}&nbsp;-&nbsp;{props.match.goalsTeam2}
      </div>
    </>
  );
};

export default TeamMatchGroup;
