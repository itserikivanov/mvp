import * as pinia from 'pinia'
import type { Mock } from 'vitest'
import type { ComputedRef } from 'vue'

type WritableGetters<Getters> = {
  [K in keyof Getters]: Getters[K] extends ComputedRef<infer T> ? T : never
}

declare type Method = (...args: any[]) => any

type MockActions<Actions> = {
  [A in keyof Actions]: Actions[A] extends Method ? Mock<Actions[A]> : Actions[A]
}

declare interface SetupStoreHelpers {
  action: <Fn extends Method>(fn: Fn) => Fn
}

declare module 'pinia' {
  // override defineStore setup store declaration
  // to make getters writable and actions mockable in unit tests
  export function defineStore<Id extends string, SS>(
    id: Id,
    storeSetup: (helpers: SetupStoreHelpers) => SS,
    options?: DefineSetupStoreOptions<
      Id,
      _ExtractStateFromSetupStore<SS>,
      _ExtractGettersFromSetupStore<SS>,
      _ExtractActionsFromSetupStore<SS>
    >,
  ): StoreDefinition<
    Id,
    _ExtractStateFromSetupStore<SS>,
    Record<string, never>,
    WritableGetters<_ExtractGettersFromSetupStore<SS>> &
      MockActions<_ExtractActionsFromSetupStore<SS>>
  >
}
