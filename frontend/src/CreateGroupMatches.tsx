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

    // If no match exists, create a new one
    if (thisMatches.length === 0) {
      const thisMatch = createEmptyMatch();
      thisMatch.groupIndex = groupIndex;
      thisMatch.matchIndex = matchIndex;
      updateCountry(thisMatch, selectedCountry);
      setMatches([...matches, thisMatch]);

    // If a match exists, update it
    } else {
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
  }

  type UpdateCountry = (match: MatchInStore, selectedCountry: string) => void
  const updateCountry1: UpdateCountry = (match, selectedCountry) => {
    match.country1 = selectedCountry
  }

  const updateCountry2: UpdateCountry = (match, selectedCountry) => {
    match.country2 = selectedCountry
  }

  return (
    <div class="m-4 p-2 border shadow">
      <h1 class="underline mb-2">{groupName}</h1>
      <For each={[0, 1, 2, 3, 4, 5]}>
        {(matchIndex) => (
          <div class="flex">
            <div class="m-4">
              <select
                name={groupName + "-" + (matchIndex + 1).toString() + "-1"}
                id={groupName + "-" + (matchIndex + 1).toString() + "-1"}
                onChange={onChangeMatchCountry(props.groupIndex, matchIndex, updateCountry1)}
              >
                <option value="" disabled selected>Select Country</option>
                <For each={countries}>
                  {(country) => (
                    <option value={country.name}>{country.name}</option>
                  )}
                </For>
              </select>
            </div>
            <div class="pt-4"> - </div>
            <div class="m-4">
              <select
                name={groupName + "-" + (matchIndex + 1).toString() + "-2"}
                id={groupName + "-" + (matchIndex + 1).toString() + "-2"}
                onChange={onChangeMatchCountry(props.groupIndex, matchIndex, updateCountry2)}
                value=""
              >
                <option value="" disabled selected>Select Country</option>
                <For each={countries}>
                  {(country) => (
                    <option value={country.name}>{country.name}</option>
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
