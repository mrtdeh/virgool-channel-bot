until test; do
    echo "GO APP EXITED : $?.  Respawning.." >&2
    sleep 1
done