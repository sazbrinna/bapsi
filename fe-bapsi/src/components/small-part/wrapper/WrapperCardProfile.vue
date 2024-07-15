<template>
  <div class="wrapper title">
    <div class="wrapper row">
      <div v-for="(user, index) in users" :key="index">
        <CardProfile v-if="!user.is_keluar" :data="user" />
      </div>
    </div>

    <span class="judul">
      <h2>Sudah pulang</h2>
    </span>

    <div class="wrapper row">
      <div v-for="(user, index) in users" :key="index">
        <CardProfile v-if="user.is_keluar" :data="user" />
      </div>
    </div>
  </div>
</template>

<script>
import CardProfile from "@/components/small-part/CardProfileHadir.vue";

export default {
  components: {
    CardProfile,
  },

  props: {
    users: Object,
  }
}
</script>

<style scoped>
.wrapper {
  gap: 1rem;
}

.title {
  gap: 2rem !important;
}

span.judul {
  border-bottom: 3px solid var(--primary);
  min-width: 300px;
  margin: 1rem 0;
}
</style>