import { defineStore } from "pinia";

export const useFileStore = defineStore("files", {
  state: () => ({ files: [] }),
});
