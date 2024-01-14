<template>
  <div :class="{ 'dark-mode': darkMode }">
    <EmailSearch />
  </div>
</template>

<script>
import EmailSearch from "./components/EmailSearch.vue";
export default {
  data() {
    return {
      darkMode: true,
    };
  },
  components: {
    EmailSearch,
  },
};
</script>
