<template>
  <div class="title-page">
    <div class="wrapper">
      <div class="title">
        <h1>{{ title }}</h1>
        <h4>{{ description }}</h4>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    description: String,
  }
}
</script>

<style scoped>
.title-page {
  display: flex;
  margin-top: 100px;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;
  width: 100%;
}

.title-page h4 {
  color: var(--dark-gray);
}

.title-page .title {
  gap: 1rem 0;
}

.title-page .wrapper {
  gap: 2rem 0;
}
</style>