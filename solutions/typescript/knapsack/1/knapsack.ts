type Item = {
  weight: number
  value: number
}

export function maximumValue({
  maximumWeight,
  items,
}: {
  maximumWeight: number
  items: Item[]
}): number {
  const itemsFiltered = items.filter((item) => item.weight <= maximumWeight);
  if (itemsFiltered.length === 0) return 0;

  const a = itemsFiltered.map((item, i) => {
    return (
      item.value + 
      maximumValue({
        maximumWeight: maximumWeight - item.weight,
        items: itemsFiltered.slice(i + 1),
      })
    );
  });

  return Math.max(...a);
}
