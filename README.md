# go-uuid
A simple and secure UUID generation script meant for Uniqueness and Efficiency

go-uuid is a tiny piece of software based on the operating system's secure random generation to provide callers with unique identifiers based on a commonly used pattern xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxxxxxx. I have myself run a simple test program to check for potential collisions and managed to get 0 collision for up to 2^32−1 uuids generated. The provided testing unit is configured to only compute 2^16−1 entries for obvious reasons.
The identifiers generated are not RFC 4122 compatible and are not meant to be any time soon. This is a pure standalone version of what one could call a unique-random key generation.
