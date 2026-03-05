import requests
import threading
import time

BASE = "http://localhost:8080"

# ---------- PARTY LOBBY ----------

party_players = ["ryan", "john", "alice", "bob"]

# ---------- SOLO LOBBY ----------

solo_players = ["mike", "sarah", "alex", "tom"]


def create_party():

    print("\nCreating party...")

    r = requests.get(f"{BASE}/party/create?player={party_players[0]}")
    party = r.json()

    party_id = party["id"]

    print("Party created:", party_id)

    # friends join
    for p in party_players[1:]:
        requests.get(f"{BASE}/party/join?party={party_id}&player={p}")
        print(p, "joined party")

    return party_id


def search_with_party(party_id):

    print("\nParty searching matchmaking...")

    r = requests.get(f"{BASE}/matchmaking/search?party={party_id}")

    try:
        print("\nParty match result:")
        print(r.json())
    except:
        print("Party waiting...")


# ---------- SOLO PLAYERS ----------

def solo_queue(player):

    r = requests.get(f"{BASE}/matchmaking/solo?player={player}")

    try:
        data = r.json()
        print(f"\nMATCH CREATED by {player}")
        print(data)
    except:
        print(player, "searching...")


# ---------- RUN SIMULATION ----------

def run_simulation():

    # create party lobby
    party_id = create_party()

    # party matchmaking
    party_thread = threading.Thread(target=search_with_party, args=(party_id,))
    party_thread.start()

    # solo matchmaking lobby
    threads = []

    for p in solo_players:
        t = threading.Thread(target=solo_queue, args=(p,))
        threads.append(t)
        t.start()
        time.sleep(0.3)

    for t in threads:
        t.join()

    party_thread.join()

    print("\nSimulation complete.")


if __name__ == "__main__":
    run_simulation()