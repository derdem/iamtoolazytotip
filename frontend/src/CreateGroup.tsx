import { Component, For, JSX } from "solid-js";
import {
  Country,
  Strength,
  getGroup,
  groupIndex,
  groups,
  matches,
  setGroupIndex,
  setGroups,
  setMatches,
} from "./groupStore";

interface CreateGroupProps {
  groupIndex: Symbol;
}

const CreateGroup: Component<CreateGroupProps> = (props) => {

  type PropertyHandler = (c: Country, ev: string) => void;
  type UpdateCountryProperty = (ci: number, ev: string, ph: PropertyHandler) => void;
  const updateCountryProperty: UpdateCountryProperty = (countryIndex, eventValue, propertyHandler) => {
    const group = {...getGroup(props.groupIndex)};

    const country = { ...group.countries[countryIndex] };
    propertyHandler(country, eventValue);
    const groupIndex = props.groupIndex;
    const countries = [...group.countries];
    countries[countryIndex] = country;

    group.countries = countries;
    const updatedGroups = groups.filter(group => group.index !== groupIndex);
    setGroups([...updatedGroups, group]);
  }

  const mutateName = (country: Country, name: string) => {
    country.name = name;
  }

  const mutateStrength = (country: Country, strength: string) => {
    country.strength = Number(strength);
  }

  type UpdateCountryName = (
    ci: number
  ) => JSX.InputEventHandlerUnion<HTMLInputElement, InputEvent>;
  const updateCountryName: UpdateCountryName =
    (countryIndex: number) => (event) => {
      updateCountryProperty(countryIndex, event.currentTarget.value, mutateName);
    };

  type UpdateCountryStrength = (
    ci: number
  ) => JSX.ChangeEventHandlerUnion<HTMLSelectElement, Event>;
  const updateCountryStrength: UpdateCountryStrength =
    (countryIndex: number) => (event) => {
      updateCountryProperty(countryIndex, event.currentTarget.value, mutateStrength);
    };

  const getCountryName = (countryIndex: number) => {
    const group = getGroup(props.groupIndex);
    return group.countries[countryIndex].name;
  };

  const getGroupName = (groupIndex: Symbol) => {
    const group = getGroup(props.groupIndex);
    return group.groupName;
  };

  const deleteGroup = () => {
    const updatedGroupIndex = groupIndex.filter(
      (oneIndex) => oneIndex !== props.groupIndex
    );
    const updatedGroups = groups.filter(
      (group) => group.index !== props.groupIndex
    );
    const updatedGroupMatches = matches.filter(
      (match) => match.groupIndex !== props.groupIndex
    );
    setMatches(updatedGroupMatches);
    setGroupIndex(updatedGroupIndex);
    setGroups(updatedGroups);
  };

  return (
    <div class="m-4 p-2 border shadow">
      <div class="flex justify-between">
        <h1 class="underline mb-2">{getGroupName(props.groupIndex)}</h1>
        <button
          class="ml-4 p-2 bg-red-500 text-white rounded-lg"
          onClick={deleteGroup}
        >
          Delete Group
        </button>
      </div>
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
                data-cy={`${getGroupName(props.groupIndex)}-${countryIndex}`}
              ></input>
            </div>
            <div class="m-4">
              <label class="p-1 bg-white left-2 -top-3 text-sm">Strength</label>
              <select
                name="Strength"
                id="strength"
                onChange={updateCountryStrength(countryIndex)}
              >
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
