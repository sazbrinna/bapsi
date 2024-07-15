<template>
  <div class="wrapper">
    <TitlePart title="Konfigurasi" description="Bapsi Gunadarma" />
    <DaftarRegion />
    <DaftarDivisi />
  </div>
</template>

<script>
import TitlePart from "@/components/TitlePart.vue";
import DaftarDivisi from "@/components/konfigurasi/daftar/DaftarDivisi.vue";
import DaftarRegion from "@/components/konfigurasi/daftar/DaftarLokasi.vue";

export default {
  components: {
    TitlePart,
    DaftarDivisi,
    DaftarRegion,
  }
}
</script>