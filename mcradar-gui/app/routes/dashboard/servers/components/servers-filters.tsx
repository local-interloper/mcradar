import { Select, Input, ListBox, Surface, Label } from "@heroui/react";
import { SERVER_TYPE_LABELS } from "~/shared/strings/server-type-labels";
import { SERVER_TYPES, type ServerType } from "~/shared/types/server-type";
import type { ServersFiltersType } from "../types/servers-filters-type";

export interface Props {
  setFilters: React.Dispatch<React.SetStateAction<ServersFiltersType>>;
}

export const ServersFilters = ({ setFilters }: Props) => {
  return (
    <Surface className="flex gap-2 p-2 rounded-lg">
      <Input
        type="text"
        variant="secondary"
        placeholder="Server Version"
        onChange={(e) =>
          setFilters((prev) => ({ ...prev, version: e.target.value }))
        }
      />
      <Select
        placeholder="Server Type"
        className="w-40"
        variant="secondary"
        selectionMode="multiple"
        onChange={(selection) =>
          setFilters((prev) => ({ ...prev, types: selection as ServerType[] }))
        }
      >
        <Select.Trigger>
          <Select.Value>
            {({ selectedItems, isPlaceholder }) =>
              isPlaceholder
                ? `Server Type`
                : `Server Type (${selectedItems.length})`
            }
          </Select.Value>
          <Select.Indicator />
        </Select.Trigger>
        <Select.Popover>
          <ListBox selectionMode="multiple">
            {SERVER_TYPES.map((serverType, i) => (
              <ListBox.Item
                key={i}
                id={serverType}
                textValue={SERVER_TYPE_LABELS[serverType]}
              >
                {SERVER_TYPE_LABELS[serverType]}
                <ListBox.ItemIndicator />
              </ListBox.Item>
            ))}
          </ListBox>
        </Select.Popover>
      </Select>
    </Surface>
  );
};
