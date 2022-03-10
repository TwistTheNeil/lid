import { defineStore } from "pinia";

export const useFileStore = defineStore("files", {
  state: () => ({ files: [] }),
  actions: {
    async refresh() {
      const response = await fetch(`/api/v1/files`);
      const responsejson = await response.json();
      this.files = responsejson;
    },
  },
});
