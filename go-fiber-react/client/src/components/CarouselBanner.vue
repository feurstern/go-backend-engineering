<script setup lang="ts">
import { onMounted } from 'vue';
import { carouselData } from './constant';
import useCarouselBanner from './CarouselBanner';

const { slide, fetchUserList } = useCarouselBanner();
onMounted(() => {
  fetchUserList()
    .catch((err) => console.log('error', err))
    .then();
});
</script>

<template>
  <div class="q-pa-md">
    <div class="q-gutter-md">
      <q-carousel
        v-model="slide"
        transition-prev="jump-right"
        transition-next="jump-left"
        swipeable
        animated
        control-color="white"
        prev-icon="arrow_left"
        next-icon="arrow_right"
        navigation-icon="radio_button_unchecked"
        navigation
        padding
        arrows
        height="300px"
        class="bg-blue text-white shadow-1 rounded-borders"
      >
        <!-- the render -->

        <q-carousel-slide
          v-for="(x, i) in carouselData"
          :key="i"
          :name="x.iconName"
          class="column no-wrap flex-center"
        >
          <q-icon name="style" size="56px" />
          <div class="q-mt-md text-center">
            {{ x.text }}
          </div>
        </q-carousel-slide>

        <!-- end  of fucking render -->
      </q-carousel>
    </div>
  </div>
</template>
