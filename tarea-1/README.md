# To run: 'go run 01.go'

Desarrolla y/o documenta una implementación apropiada para las siguientes clases:

- STACK (lifo)
- QUEUE (fifo)
- TABLE/HASH/DICTIONARY (order) (las puedes implementar “desde 0” o usar alguna librería “pública”)

Las clases deben contener métodos para soportar las principales operaciones de acceso y manipulación (clásicas). NO
tienen que ser clases, si el lenguaje de desarrollo no las soporta.

## Casos de Prueba STACK (push, pop, top, isEmpty)

1. Push 1 --> [1] --> Top debe ser 1
2. Push 2 --> [1, 2] --> Top debe ser 2
3. Push 3 --> [1, 2, 3] --> Top debe se 3
4. Pop --> [1, 2] --> Retorna 3
5. Top --> [1, 2] --> Debe ser 2
6. IsEmpty --> [1, 2] --> Debe ser false
7. Pop --> [1] --> Retorna 2
8. Pop --> [] --> Retorna 1
9. IsEmpty --> Debe ser true

## Casos de Prueba QUEUE (enqueue, dequeue, head, tail, isEmpty)

1. Enqueue --> [1] --> Head should be 1, Tail should be 1
2. Enqueue --> [1, 2] --> Head should be 1, Tail should be 2
3. Enqueue --> [1, 2, 3] --> Head should be 1, Tail should be 3
4. Dequeue --> [2, 3] --> Should return 1
5. Head --> [2, 3] --> Should be 2
6. Dequeue --> [3] --> Should return 2
7. Head --> [3] --> Should be 3
8. Tail --> [3] --> Should be 3
9. Enqueue --> [3, 4] --> Tail should be 4
10. IsEmpty --> [3, 4] --> Should be false
11. 2 Dequeue --> [] --> Length should be 0
12. IsEmpty --> [] --> Should be true

## CASOS de Prueba DICTIONARY (add, access, update, delete, length)

1. Add --> {"Diego": 22}
2. Add --> {"Diego": 22, "Roberto": 21}
3. Access "Roberto" --> Should return 21
4. Length --> Should be 2
5. Update --> {"Diego": 22, "Roberto": 32}
6. Access "Roberto" --> Should return 32
7. Delete "Roberto" --> {"Diego": 22}
8. Length should be 1
9. Add --> {"Diego": 22, "Jose": 20}
10. Access "Jose" --> Should return 20
11. Length --> Should be 2
12. Access "Diego" --> Should be 22
13. Delete "Diego --> Length should be 1
14. Delete "Jose --> Length should be 0
