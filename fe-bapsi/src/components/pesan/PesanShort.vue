<template>
  <div class="pesan">
    <div class="wrapper">
      <div class="title row">
        <div class="title-name">
          <h3>Pesan</h3>
          <p class="gray">Buat pengumuman atau pesan untuk asisten lain</p>
        </div>
        <button class="bt-1">+ Pesan Baru</button>
      </div>

      <div class="wrapper-card-pesan row">

        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        <CardPesan />
        
      </div>
    </div>

  </div>
</template>

<script>
import CardPesan from "@/components/pesan/CardPesan.vue";

export default {
  components: {
    CardPesan
  }
}
</script>

<style scoped>
.wrapper-card-pesan {
  width: 100%;
  gap: 1rem;
  align-items: center;
  justify-content: center;
}

.pesan {
  width: 100%;
}

.pesan .wrapper {
  gap: 1rem;
}

.pesan .title h3 {
  text-align: start;
}

.pesan .title {
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.pesan .title-name {
  align-items: flex-start;
}
</style>