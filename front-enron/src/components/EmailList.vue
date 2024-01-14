<template>
  <div class="email-list-container">
    <table class="table-auto w-full">
      <thead>
        <tr>
          <th class="border px-4 py-2 bg-gray-900 text-blue-300">From</th>
          <th class="border px-4 py-2 bg-gray-900 text-blue-300">To</th>
          <th class="border px-4 py-2 bg-gray-900 text-blue-300">Subject</th>
          <th class="border px-4 py-2 bg-gray-900 text-blue-300">Content</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="email in paginatedEmails"
          :key="email.id"
          @click="selectEmail(email)"
          :class="{'selected-row': email === selectedEmail}"
          style="cursor: pointer;"
        >
          <td class="border px-4 py-2">{{ truncateText(email.from, 15) }}</td>
          <td class="border px-4 py-2">{{ truncateText(email.to, 15) }}</td>
          <td class="border px-4 py-2">{{ truncateText(email.subject, 15) }}</td>
          <td class="border px-4 py-2 overflow-wrap ">{{ truncateText(email.content, 150) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
  
  <!-- Pagination -->
  <div class="pagination-container flex justify-center mt-4">
    <button @click="goToPreviousPage" :disabled="currentPage === 1" class="mx-2 px-3 py-1 bg-blue-300 text-white rounded hover:bg-blue-500" v-if="currentPage > 1">Previous page</button>
    <span class="mx-2 px-3 py-1">{{ currentPage }} / {{ totalPages }}</span>
    <button @click="goToNextPage" :disabled="currentPage === totalPages" class="mx-2 px-3 py-1 bg-blue-300 text-white rounded hover:bg-blue-500" v-if="currentPage < totalPages">Next Page</button>
  </div>
</template>

<script>
export default {
  props: {
    emails: {
      type: Array,
      default: () => [],
    },
  },
  data() {
    return {
      currentPage: 1,
      selectedEmail: null,
    };
  },
  methods: {
    truncateText(text, maxLength) {
      if (text.length > maxLength) {
        return text.slice(0, maxLength) + '...';
      }
      return text;
    },
    goToPreviousPage() {
      if (this.currentPage > 1) {
        this.currentPage -= 1;
      }
    },
    goToNextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage += 1;
      }
    },
    selectEmail(email) {
      this.selectedEmail = email;
      this.$emit("email-selected", email);
    },
  },
  computed: {
    totalPages() {
      return Math.ceil(this.emails.length / 10); 
    },
    paginatedEmails() {
      const startIndex = (this.currentPage - 1) * 10;
      const endIndex = startIndex + 10;
      return this.emails.slice(startIndex, endIndex);
    },
  },
};
</script>

<style>
.email-list-container {
  width: 90%;
  margin: 0 auto;
}

.pagination-container {
  margin-top: 20px;
}
.selected-row {
  background-color: #202020; 
  color: #4CAF50;
}
</style>
