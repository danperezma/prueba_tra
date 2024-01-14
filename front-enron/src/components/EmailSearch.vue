<template>
  <div class="h-screen flex flex-col items-center justify-center">
    <h1 class="text-3xl font-semibold mb-4">Email Search Engine</h1>
    <div class="w-3/4 mb-4">
      <div class="relative flex">
        <input
          v-model="searchTerm"
          @keyup.enter="search"
          placeholder="Search by keyword"
          autofocus="true"
          class="p-2 border rounded-l-md w-full"
          :style="{ color: 'black'}"
        />
        <button @click="search" class="bg-blue-500 text-white p-2 rounded-r-md hover:bg-blue-700 absolute right-0 top-0 h-full">
          <i class="fas fa-search"></i>
        </button>
      </div>
    </div>
    <div class="overflow-y-auto">
      <div v-if="emails === null || emails === undefined" class="text-center my-4">
        <p class="text-xl text-red-300 font-bold">No match found please try with another keyword.</p>
        <img src="../assets/logo.png" alt="Logo" class="mt-4 mx-auto w-80" />
      </div>
      <div v-else-if="emails.length === 0" class="text-center my-4">
        <p class="text-xl text-white-700">¡Make a search!</p>
        <img src="../assets/logo.png" alt="Logo" class="mt-4 w-80" />
      </div>
      <div v-else>
        <div class="flex">
          <div class="w-3/4" style="overflow-x: auto;">
            <email-list :emails="emails" @email-selected="displayEmail"/>
          </div>
          <div class="w-1/4">
            <email-view :selectedEmail="selectedEmail" :searchTerm="searchTermProp"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import EmailList from "./EmailList.vue";
import EmailView from "./EmailView.vue";

export default {
  components: {
    EmailList,
    EmailView,
  },
  data() {
    return {
      searchTerm: "",
      emails: [], 
      selectedEmail: null, 
      searchTermProp : null,
    };
  },
  methods: {
    async search() {
      try {
        const requestBody = { searchTerm: this.searchTerm };
        const response = await fetch("http://localhost:8080/api/query", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(requestBody),
        });
        if (!response.ok) {
          throw new Error("The mailing list was not available.");
        }
        this.emails = await response.json();
        this.selectedEmail = null;
        this.searchTermProp = this.searchTerm;
      } catch (error) {
        console.error(error);
        this.emails = null;
      }
    },
    displayEmail(email) {
      this.selectedEmail = email;
    },
  },
};
</script>

<style>
</style>
