import { Component } from "solid-js";
import { A } from "@solidjs/router";

interface EnterTournamentCard {
  name: string;
  url: string;
}

const EnterTournamentCard: Component<EnterTournamentCard> = (props) => {
  return (
    <A href={props.url} class="no-underline">
      <div class="ring-2 ring-sky-800 w-60">
        <div class="w-100 bg-sky-800 text-white text-center p-4">
          {props.name}
        </div>
        <div class="text-center h-40 py-auto"></div>
      </div>
    </A>
  );
};

export default EnterTournamentCard;
