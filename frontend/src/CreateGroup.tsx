import { Component, For, JSX } from "solid-js";
import { Strength, groups, setGroups } from "./groupStore";

interface CreateGroupProps {
  groupIndex: number;
}

const CreateGroup: Component<CreateGroupProps> = (props) => {
  console.log("CreateGroup rendered")
  const updateCountryName: (
    ci: number
  ) => JSX.InputEventHandlerUnion<HTMLInputElement, InputEvent> =
    (countryIndex: number) => (event) => {
      const countryName = event.currentTarget.value;
      const country = {...groups[props.groupIndex].countries[countryIndex]};
      country.name = countryName;

      const groupIndex = props.groupIndex;
      const countries = [...groups[props.groupIndex].countries];
      countries[countryIndex] = country;

      const group = {...groups[props.groupIndex]};
      group.countries = countries;

      const updatedGroups = [...groups];
      updatedGroups[groupIndex] = group;

      setGroups(updatedGroups);
    };

  const updateCountryStrength: (
    ci: number
  ) => JSX.ChangeEventHandlerUnion<HTMLSelectElement, Event> =
    (countryIndex: number) => (event) => {
      const countryStrength = Number(event.currentTarget.value);
      const country = {...groups[props.groupIndex].countries[countryIndex]};
      country.strength = countryStrength;

      const groupIndex = props.groupIndex;
      const countries = [...groups[props.groupIndex].countries];
      countries[countryIndex] = country;

      const group = {...groups[props.groupIndex]};
      group.countries = countries;

      const updatedGroups = [...groups];
      updatedGroups[groupIndex] = group;

      setGroups(updatedGroups);
    };

  const getCountryName = (countryIndex: number) => {
    return groups[props.groupIndex].countries[countryIndex].name;
  };

  return (
    <div class="m-4 p-2 border shadow">
      <h1 class="underline mb-2">{groups[props.groupIndex].groupName}</h1>
      <For each={[0, 1, 2, 3]}>
        {(countryIndex) => (
          <div class="flex">
            <div class="relative m-4">
              <label class="p-1 absolute bg-white left-2 -top-3 text-sm">
                Country
              </label>
              <input
                name="group-name"
                value={getCountryName(countryIndex)}
                class="p-4 mr-2 outline-1 border-2 rounded-lg outline-slate-200 focus:outline-slate-400 shadow"
                onInput={updateCountryName(countryIndex)}
              ></input>
            </div>
            <div class="m-4">
              <label class="p-1 bg-white left-2 -top-3 text-sm">
                Strength
              </label>
              <select name="Strength" id="strength" onChange={updateCountryStrength(countryIndex)}>
                <option value={Strength.Weak}>{Strength.Weak}</option>
                <option value={Strength.Medium}>{Strength.Medium}</option>
                <option value={Strength.Strong}>{Strength.Strong}</option>
              </select>
            </div>
          </div>
        )}
      </For>
    </div>
  );
};

export default CreateGroup;
