module Test.Main where

import Prelude

import Effect (Effect)
import Effect.Console (log)
import Node.Process (getEnv, cwd, pid, version, execPath, platform)
import Test.Assert (assert)
import Data.Maybe (isJust)
import Foreign.Object as Object

main :: Effect Unit
main = do
  env <- getEnv
  -- Basic sanity checks
  assert $ Object.size env > 0
  
  c <- cwd
  assert $ c /= ""
  
  -- Evaluate pid
  let _ = pid
  
  let _ = version
  
  exec <- execPath
  assert $ exec /= ""
  
  assert $ isJust platform
  
  log "Tests passed 🎉"
