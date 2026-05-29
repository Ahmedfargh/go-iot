<template>
  <div class="space-y-lg">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end gap-4">
      <div>
        <h1 class="font-headline-md text-headline-md text-on-surface">System Overview</h1>
        <p class="font-body-md text-body-md text-on-surface-variant mt-xs">Real-time monitoring and high-level health metrics.</p>
      </div>
      <button class="bg-primary-container hover:bg-primary-fixed text-on-primary-container font-label-caps text-label-caps px-lg py-sm rounded-lg transition-all duration-200 flex items-center gap-sm shadow-sm">
        <span class="material-symbols-outlined">refresh</span>
        REFRESH
      </button>
    </div>

    <!-- Quick Stats -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-md">
      <div class="glass-card p-md rounded-xl">
        <div class="flex items-center gap-xs mb-2">
          <span class="material-symbols-outlined text-primary">dns</span>
          <p class="font-label-caps text-label-caps text-on-surface-variant uppercase">Server Load</p>
        </div>
        <div class="flex items-end justify-between mt-xs">
          <span class="font-display-lg text-display-lg text-on-surface">42%</span>
          <span class="text-primary font-data-md text-data-md bg-primary/10 px-2 rounded-md">+5%</span>
        </div>
      </div>
      <div class="glass-card p-md rounded-xl">
        <div class="flex items-center gap-xs mb-2">
          <span class="material-symbols-outlined text-secondary">router</span>
          <p class="font-label-caps text-label-caps text-on-surface-variant uppercase">Network Traffic</p>
        </div>
        <div class="flex items-end justify-between mt-xs">
          <span class="font-display-lg text-display-lg text-on-surface">1.2<span class="text-lg">gbps</span></span>
          <span class="text-secondary font-data-md text-data-md bg-secondary/10 px-2 rounded-md">Stable</span>
        </div>
      </div>
      <div class="glass-card p-md rounded-xl ring-1 ring-white/10 relative overflow-hidden">
        <div class="shimmer absolute inset-0"></div>
        <div class="relative z-10">
            <div class="flex items-center gap-xs mb-2">
            <span class="material-symbols-outlined text-tertiary">warning</span>
            <p class="font-label-caps text-label-caps text-on-surface-variant uppercase">Active Alerts</p>
            </div>
            <div class="flex items-end justify-between mt-xs">
            <span class="font-display-lg text-display-lg text-tertiary">3</span>
            <button class="text-on-surface font-label-caps text-[10px] underline hover:text-tertiary transition-colors">View All</button>
            </div>
        </div>
      </div>
    </div>
    
    <div class="flex items-center justify-center h-64 border border-dashed border-white/10 rounded-xl bg-surface-container/20">
      <p class="font-body-md text-on-surface-variant/40 italic">Overview charts and activity feeds would go here...</p>
    </div>
  </div>
</template>

<script setup>
</script>
