<template>
  <section class="full">
    <h3>Latitude : {{ latitude }}</h3>
    <h3>Longitude : {{ longitude }}</h3>
  </section>
</template>

<script>
export default {
  data() {
    return {
      latitude: "",
      longitude: "",
    }
  },

  created() {
    if ("geolocation" in navigator) {
      // get user geolocation information

      navigator.geolocation.getCurrentPosition(
        position => {
          // Get latitude from the position object
          this.latitude = position.coords.latitude;
          this.longitude = position.coords.longitude;
        });
    }
  }
}
</script>