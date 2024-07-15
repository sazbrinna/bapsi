<template>
  <div class="card-profile row">
    <div class="image">
      <img src="@/assets/icons/blankProfile.jpg" alt="">
    </div>
    <div class="data">
      <h3 @click.prevent="handlerProfile(data.id_user)">{{ data.nama }}</h3>
      <p v-if="data.jam_pulang != '00:00'"><strong>{{ data.lokasi_absen }}</strong> | {{ data.jam_absen }} | {{ data.jam_pulang }}</p>
      <p v-else><strong>{{ data.lokasi_absen }}</strong> | {{ data.jam_absen }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    data: Object,
  },

  methods: {
    handlerProfile(path) {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });

      this.$router.push({ path: "/profile/" + path });
    }
  }
}
</script>

<style scoped>
.card-profile {
  padding: 15px;
  align-items: center;
  gap: 1rem;
  border-style: solid;
  border-color: var(--secondary);
  border-radius: 8px;
  width: 360px;
  max-width: 100%;
  box-sizing: border-box;
}

.data {
  text-align: left;
}

.card-profile .image {
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-profile img {
  width: 75px;
  height: 75px;
  object-fit: cover;
  border-radius: 100%;
}

.card-profile h3 {
  color: var(--primary);
  cursor: pointer;
}
</style>