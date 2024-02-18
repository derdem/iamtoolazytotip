import { Component, For, JSX } from "solid-js";
import { MatchInStore, createEmptyMatch, groups, matches, setMatches } from "./groupStore";

interface CreateGroupProps {
  groupIndex: number;
}

const CreateGroupMatches: Component<CreateGroupProps> = (props) => {
  const groupName = groups[props.groupIndex].groupName;
  const countries = groups[props.groupIndex].countries;

  type OnChangeMatch = (gi: number, mi: number, uc: UpdateCountry) => JSX.ChangeEventHandlerUnion<HTMLSelectElement, Event>
  const onChangeMatchCountry: OnChangeMatch = (groupIndex: number, matchIndex: number, updateCountry: UpdateCountry) => (event) => {
    const selectedCountry = event.currentTarget.value;
    const thisMatches = matches.filter((match) => {
      return match.groupIndex === groupIndex && match.matchIndex === matchIndex;
    })

    if (thisMatches.length !== 1) {
      throw new Error("Match should exist")
    }

    const thisMatch = thisMatches[0];
    const thisMatchNew = createEmptyMatch();
    thisMatchNew.groupIndex = groupIndex;
    thisMatchNew.matchIndex = matchIndex;
    thisMatchNew.country1 = thisMatch.country1;
    thisMatchNew.country2 = thisMatch.country2;
    updateCountry(thisMatchNew, selectedCountry);
    const allOtherMatches = matches.filter((match) => {
      return match.groupIndex !== groupIndex || match.matchIndex !== matchIndex;
    })
    setMatches([...allOtherMatches, thisMatchNew]);

  }

  const groupMatches = matches.filter((match) => {
    return match.groupIndex === props.groupIndex;
  });

  type UpdateCountry = (match: MatchInStore, selectedCountry: string) => void
  const updateCountry1: UpdateCountry = (match, selectedCountry) => {
    match.country1 = selectedCountry
  }

  const updateCountry2: UpdateCountry = (match, selectedCountry) => {
    match.country2 = selectedCountry
  }

  return (
    <div class="m-4 p-2 border shadow">
      <h1 class="underline mb-2">Matches for: {groupName}</h1>
      <For each={groupMatches}>
        {(match) => (
          <div class="flex">
            <div class="m-4">
              <select
                name={groupName + "-" + (match.matchIndex + 1).toString() + "-1"}
                id={groupName + "-" + (match.matchIndex + 1).toString() + "-1"}
                onChange={onChangeMatchCountry(props.groupIndex, match.matchIndex, updateCountry1)}
              >
                <option value="" disabled selected={match.country1 === ""}>Select Country</option>
                <For each={countries}>
                  {(country) => (
                    <option value={country.name} selected={country.name === match.country1}>{country.name}</option>
                  )}
                </For>
              </select>
            </div>
            <div class="pt-4"> - </div>
            <div class="m-4">
              <select
                name={groupName + "-" + (match.matchIndex + 1).toString() + "-2"}
                id={groupName + "-" + (match.matchIndex + 1).toString() + "-2"}
                onChange={onChangeMatchCountry(props.groupIndex, match.matchIndex, updateCountry2)}
                value=""
              >
                <option value="" disabled selected={match.country2 === ""}>Select Country</option>
                <For each={countries}>
                  {(country) => (
                    <option value={country.name} selected={country.name === match.country2}>{country.name}</option>
                  )}
                </For>
              </select>
            </div>
          </div>
        )}
      </For>
    </div>
  );
}

export default CreateGroupMatches;
