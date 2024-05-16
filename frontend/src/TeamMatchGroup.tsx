import { Component, Show } from "solid-js";

const TeamMatchGroup: Component<{[key: string]: any}> = (props) => {
  return (
    <>
      <div class="">{props.matchResult.match.team1.name}</div>
      <div class="px-4">&nbsp;-&nbsp;</div>
      <div class="">{props.matchResult.match.team2.name}</div>
      <div class="flex-grow pl-4">
        :&nbsp;&nbsp;&nbsp;
        {props.matchResult.team1Goals}&nbsp;-&nbsp;{props.matchResult.team2Goals}
      </div>
    </>
  );
};

export default TeamMatchGroup;
