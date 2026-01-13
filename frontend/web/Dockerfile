FROM node:24-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
# Node 25 doesn't include corepack anymore
RUN corepack enable

WORKDIR /app/web

FROM base AS development
CMD ["scripts/start.sh"]