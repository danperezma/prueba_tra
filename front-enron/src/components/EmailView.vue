<template>
  <div>
    <div v-if="selectedEmail === null || selectedEmail === undefined" class="text-center my-4">
      <p class="text-white text-xl">Select an e-mail to view its content.</p>
    </div>
    <div v-else>
      <div class="max-h-[85vh] overflow-y-auto">
        <h1 class="text-3xl font-semibold">{{ selectedEmail.subject }}</h1>
        <h2 class="text-xl text-green-300">From: {{ selectedEmail.from }}</h2>
        <h2 class="text-xl text-violet-300">To: {{ selectedEmail.to }}</h2>
        <pre class="overflow-wrap whitespace-pre-line">
          <span v-html="highlightSearchTerm(escapeHtml(selectedEmail.full_file))"></span>
        </pre>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    selectedEmail: Object,
    searchTerm: String,
  },
  methods: {
    highlightSearchTerm(fullFile) {
      const regex = new RegExp(this.searchTerm, 'gi');
      const highlightedContent = fullFile.replace(regex, `<span class="highlight">${this.searchTerm}</span>`);
      return highlightedContent;
    },
    escapeHtml(text) {
      return text.replace(/</g, "&lt;").replace(/>/g, "&gt;");
    },
  },
};
</script>

<style>
.highlight {
  background-color: yellow;
  color: black;
}
</style>
