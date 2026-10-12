import { defineConfig } from 'tsup'

export default defineConfig([
  {
    entry: { app: 'web/src/main.ts' },
    outDir: 'web/dist',
    format: ['esm'],
    platform: 'browser',
    clean: true,
  },
  {
    entry: { verify: 'verify/verify.ts' },
    outDir: 'verify/dist',
    format: ['esm'],
    platform: 'node',
    clean: true,
  },
])
