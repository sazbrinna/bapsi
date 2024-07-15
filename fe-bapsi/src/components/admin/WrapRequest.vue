<template>
<div class="request" v-if="requests">
  <h2 class="title">Request Account</h2>
  <h3><strong class="primary">{{ requests.length }}</strong> Permintaan</h3>
  <div class="wrapRequest">
    <CardRequest v-for="(request, index) in requests" :key="index" :data="request" />
  </div>
  </div>
</template>

<script>
import CardRequest from "@/components/admin/cards/RequestAccount.vue";

export default {
  props: {
    requests: [],
  },

  components: {
    CardRequest,
  }
}
</script>

<style scoped>
  .request {
    display: flex;
    width: 100%;
    align-items: center;
    justify-content: center;
    flex-direction: column;
    flex-wrap: wrap;
    gap: 1rem;
  }

  .wrapRequest {
    display: flex;
    width: 100%;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    flex-wrap: wrap;
    gap: 1rem;
  }

  h3 strong {
    font-size: 1.5rem;
  }
</style>