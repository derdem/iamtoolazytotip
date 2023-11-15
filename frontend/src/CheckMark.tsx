import { Component } from "solid-js";
import styles from "./CheckMark.module.css";

const CheckMark: Component = () => {
  return (
    <svg
      class={styles.checkmark}
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 40 40"
    >
      <path
        fill="none"
        d="M14.1 27.2l7.1 7.2 16.7-16.8"
      />
    </svg>
  );
};

export default CheckMark;
