import { FireballData } from "@/lib/types";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

const FireballTable = (props: { data: FireballData[] }) => {
  return (
    <Table>
      <TableCaption>Fireball Data</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">ID</TableHead>
          <TableHead className="text-right">Radiated Energy</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {props.data.map((fireball) => (
          <TableRow>
            <TableCell className="font-medium">{fireball.id}</TableCell>
            <TableCell className="text-right">
              {fireball.total_radiated_energy}
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default FireballTable;
