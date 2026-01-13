FROM node:24-slim AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
# Node 25 doesn't include corepack anymore
RUN corepack enable

WORKDIR /app

FROM base AS development
# svelte.config.js is needed for .svelte-kit generation
# Vite forced re-optimization is fine
COPY package.json pnpm-lock.yaml svelte.config.js /app
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
COPY . /app
CMD ["pnpm", "run", "dev", "--host"]